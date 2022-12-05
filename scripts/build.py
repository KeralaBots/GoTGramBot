import re
from typing import Union

import requests
from pathlib import Path


API_PATH = Path('scripts/api')
TEMPLATE = Path('scripts/template')
RESULTS = Path("")

TG_CORE_TYPES = {
    "String": 'string',
    "Boolean": 'bool',
    "Integer": 'int64',
    "Float": 'float64'
}

CORE_TYPES = ['int64', 'float64', 'bool', 'string']
MARKUP = ['InlineKeyboardMarkup', 'ReplyKeyboardMarkup', 'ReplyKeyboardRemove', 'ForceReply']

ARRAY_TYPE = []
ARRAY_OF_ARRAY_TYPE = []

SUBCLASS_DICT = {}

type_temp = open(TEMPLATE / 'types_common.tmpl', mode='r').read()
array_temp = open(TEMPLATE / 'array.tmpl', mode='r').read()
array_of_array_temp = open(TEMPLATE / 'array_of_array.tmpl', mode='r').read()
method_temp = open(TEMPLATE / 'methods_common.tmpl', mode='r').read()

subclass_temp = """func (v {class_name}) Get{method}() {class_name} {{
    return v
}}
"""

returned_temp = """
    var res {extra}{returned}
    return res, json.Unmarshal(r, &res) 
"""

markup_temp = """
func (m {classname}) replyMarkup() {{}}
"""

markup_content = '''
type ReplyMarkup interface {
    replyMarkup()
}
'''

optional_params_temp = """
        if opts.{name} != nil {{
            bs, err := json.Marshal(opts.{name})
            if err != nil {{
                return {is_bool}, fmt.Errorf("failed to marshal field {field_name}: %w", err)
            }}
            params["{field_name}"] = string(bs)
        }}

"""

req_params_temp = """
    if {name} != nil {{
        bs, err := json.Marshal({name})
        if err != nil {{
            return {is_bool}, fmt.Errorf("failed to marshal field {field_name}: %w", err)
        }}
        params["{field_name}"] = string(bs)
    }}

"""

opt_input_temp = """
        if opts.{name} != nil {{
            switch f := opts.{name}.(type) {{
            case string:
                _, err := os.Stat(f)
                if err != nil {{
                    params["{field_name}"] = f
                }} else {{
                    params["{field_name}"] = "attach://{field_name}"
                    data_params["{field_name}"] = f
                }}
            default:
                return {is_bool}, fmt.Errorf("unknown type for InputFile: %T", opts.{name})
            }}
        }}
"""

req_input_temp = """
    if {name} != nil {{
        switch f := {name}.(type) {{
        case string:
            _, err := os.Stat(f)
            if err != nil {{
                params["{field_name}"] = f
            }} else {{
                params["{field_name}"] = "attach://{field_name}"
                data_params["{field_name}"] = f
            }}
        default:
            return {is_bool}, fmt.Errorf("unknown type for InputFile: %T", {name})
        }}
    }}
"""

nofield_method_temp = """
{comments}
func (b *Bot) {struct_name}({fields}) ({returns}error) {{
    params := map[string]string{{}}
    data_params := map[string]string{{}}
{params}
    r, err := b.Request("{method_name}", params, data_params)
    if err != nil {{
        return {is_bool}, err
    }}
    {single_returns}
}}
"""


def snake(s: str):
    # https://stackoverflow.com/q/1175208
    s = re.sub(r"(.)([A-Z][a-z]+)", r"\1_\2", s)
    return re.sub(r"([a-z0-9])([A-Z])", r"\1_\2", s).lower()


def camel(s: str):
    return "".join([i[0].upper() + i[1:] for i in s.split("_")])


def generate_api():
    req = requests.get("https://raw.githubusercontent.com/PaulSonOfLars/telegram-bot-api-spec/main/api.json")
    if req.ok:
        return req.json()
    else:
        return None


def get_type(types):
    def_types = TG_CORE_TYPES.get(types) if TG_CORE_TYPES.get(types) is not None else types
    if def_types.startswith("Array of Array"):
        if def_types[18:] not in ARRAY_OF_ARRAY_TYPE and TG_CORE_TYPES.get(def_types[18:]) is None:
            ARRAY_OF_ARRAY_TYPE.append(def_types[18:])
        def_types = f"[][]{TG_CORE_TYPES.get(f'{def_types[18:]}')}" if TG_CORE_TYPES.get(
            f'{def_types[18:]}') is not None else f'[][]{def_types[18:]}'
    elif def_types.startswith("Array of"):
        if def_types[9:] not in ARRAY_TYPE and TG_CORE_TYPES.get(def_types[9:]) is None:
            ARRAY_TYPE.append(def_types[9:])
        def_types = f"[]{TG_CORE_TYPES.get(f'{def_types[9:]}')}" if TG_CORE_TYPES.get(
            f'{def_types[9:]}') is not None else f'[]{def_types[9:]}'
    else:
        def_types = def_types if def_types in CORE_TYPES else f'*{def_types}'
    return def_types


def get_field_type(types):
    def_types = TG_CORE_TYPES.get(types) if TG_CORE_TYPES.get(types) is not None else types
    if def_types.startswith("Array of Array"):
        if def_types[18:] not in ARRAY_OF_ARRAY_TYPE and TG_CORE_TYPES.get(def_types[18:]) is None:
            ARRAY_OF_ARRAY_TYPE.append(def_types[18:])
        if TG_CORE_TYPES.get(f'{def_types[18:]}') is not None:
            return f"{TG_CORE_TYPES.get(f'{def_types[18:]}')}", '[][]'
        else:
            return f'types.{def_types[18:]}', '[][]'
    elif def_types.startswith("Array of"):
        if def_types[9:] not in ARRAY_TYPE and TG_CORE_TYPES.get(def_types[9:]) is None:
            ARRAY_TYPE.append(def_types[9:])
        if TG_CORE_TYPES.get(f'{def_types[9:]}') is not None:
            return f"{TG_CORE_TYPES.get(f'{def_types[9:]}')}", '[]'
        else:
            return f'types.{def_types[9:]}', '[]'
    else:
        if def_types in CORE_TYPES:
            return def_types, ''
        else:
            return f'types.{def_types}', '*'


def get_field_text(field_name, def_types, field, comments=True):
    if field.get('required'):
        empty = ""
    else:
        empty = ",omitempty"
    if comments:
        text = f'    // {field.get("description")}\n' \
               f'    {camel(field_name)} {def_types} `json:"{field_name}{empty}"`\n'
    else:
        text = f'    {camel(field_name)} {def_types} `json:"{field_name}{empty}"`\n'
    return text


def get_inheritance(subclasses, subclass_dict, schema, name):
    field_text = ''
    sub_fields_list = []
    for subclass in subclasses:
        subclass_dict.update({subclass: name})
        sub_schema = schema.get(subclass)
        text = ''
        for sub_fields in sub_schema.get('fields'):
            sub_field_name = sub_fields.get('name')
            if sub_field_name in sub_fields_list:
                continue
            else:
                sub_fields_list.append(sub_field_name)
                for sub_types in sub_fields.get('types'):
                    sub_def_types = get_type(sub_types)
                    if "InputFile" in sub_def_types:
                        continue
                    if sub_field_name == "chat_id" and sub_def_types == "string":
                        continue

                    text += get_field_text(sub_field_name, sub_def_types, sub_fields, comments=False)
        field_text += text
    return field_text


def get_unmarshals():
    content = ''
    for array in ARRAY_TYPE:
        content += array_temp.format(class_name=array)

    for bi_array in ARRAY_OF_ARRAY_TYPE:
        content += array_of_array_temp.format(class_name=bi_array)
    return content


def write_types(content):
    with open(RESULTS / "types/types.go", "w+", encoding="utf-8") as type_file:
        type_file.write(
            type_temp.format(
                content=content
            )
        )


def write_methods(content):
    with open(RESULTS / "methods.go", "w+", encoding="utf-8") as method_file:
        method_file.write(
            method_temp.format(
                content=content
            )
        )


def method_fields(fields, method_name):
    required_field_list = []
    field_list_r = []
    field_list_o = []
    optional_field_list = []
    for field in fields:
        field_name = camel(field.get("name"))
        field_name = field_name[0].lower() + field_name[1:]
        field_type_text = ''
        types = field.get('types')
        if len(types) > 1:
            if types == MARKUP:
                field_type_text += 'types.ReplyMarkup'
            else:
                field_type, extra = get_field_type(types[0])
                field_type_text += extra + field_type
        else:
            field_type, extra = get_field_type(types[0])
            field_type_text += extra + field_type

        if "InputFile" in field_type_text:
            field_type_text = field_type_text.replace("*", "")
        if field.get("required"):
            data_form = {'field': field.get('name'), 'type': field_type_text}
            field_list_r.append(data_form)
            field_text = field_name + " " + field_type_text
            required_field_list.append(field_text)
        else:
            field_text = camel(field_name) + ' ' + field_type_text + f' `json:"{snake(field_name)},omitempty"`'
            optional_field_list.append(field_text)
            data_form = {'field': field.get('name'), 'type': field_type_text}
            field_list_o.append(data_form)
    if len(optional_field_list) > 0:
        required_field_list.append(f'opts *{camel(method_name)}Opts')
    method_params = ", ".join(required_field_list)
    return method_params, optional_field_list, field_list_r, field_list_o


def construct_opts(struct_list):
    r = "\n    ".join(struct_list)
    return r


def construct_params(params, optionals, is_bool):
    res = ""
    for p in params:
        new_p = p['field']
        field_name = camel(new_p)
        new_name = field_name[0].lower() + field_name[1:]
        res += format_params(new_name, p['type'], new_p, True,
                             is_bool)  # f'    params["{new_p}"] = {field_name[0].lower()}{field_name[1:]}\n'
    if len(optionals) > 0:
        res += '    if opts != nil {\n'
        for o in optionals:
            new_o = o['field']
            field_name = camel(new_o)
            res += format_params(field_name, o['type'], new_o, False,
                                 is_bool)  # f'        params["{new_o}"] = opts.{field_name}\n'
        res += '    }\n'
    return res


def construct_returns(data, returns):
    return_text = ''
    data_type, extra = get_field_type(data)
    if extra == '*[][]':
        if returns in ARRAY_OF_ARRAY_TYPE:
            return_text = f'return types.Unmarshal{returns}ArrayOfArray(r)'
        else:
            return_text = returned_temp.format(returned=data_type, extra=extra)
    elif extra == '*[]':
        if returns in ARRAY_TYPE:
            return_text = f'return types.Unmarshal{returns}Array(r)'
        else:
            return_text = returned_temp.format(returned=data_type, extra=extra)
    else:
        return_text = returned_temp.format(returned=data_type, extra=extra)
    return return_text


def format_params(raw_data, typed, param_name, required: bool = False, is_bool=False):
    if typed == 'string':
        if required:
            data = f'    params["{param_name}"] = {raw_data}\n'
        else:
            data = f'        params["{param_name}"] = opts.{raw_data}\n'
    elif typed == 'int64':
        if required:
            data = f'    params["{param_name}"] = strconv.FormatInt({raw_data}, 10)\n'
        else:
            data = f'        params["{param_name}"] = strconv.FormatInt(opts.{raw_data}, 10)\n'
    elif typed == 'float64':
        if required:
            data = f'    params["{param_name}"] = strconv.FormatFloat({raw_data}, \'E\', -1, 64)\n'
        else:
            data = f'        params["{param_name}"] = strconv.FormatFloat(opts.{raw_data}, \'E\', -1, 64)\n'
    elif typed == 'bool':
        if required:
            data = f'    params["{param_name}"] = strconv.FormatBool({raw_data})\n'
        else:
            data = f'        params["{param_name}"] = strconv.FormatBool(opts.{raw_data})\n'
    elif "types.InputFile" in typed:
        if required:
            data = req_input_temp.format(
                name=raw_data,
                field_name=param_name,
                is_bool=is_bool
            )
        else:
            data = opt_input_temp.format(
                name=raw_data,
                field_name=param_name,
                is_bool=is_bool
            )
    else:
        if required:
            data = req_params_temp.format(
                name=raw_data,
                field_name=param_name,
                is_bool=is_bool
            )
        else:
            data = optional_params_temp.format(
                name=raw_data,
                field_name=param_name,
                is_bool=is_bool
            )
    return data


def build_types(api_content: Union[dict, None]):
    content = ''
    if api_content is not None:
        content_temp = open(TEMPLATE / 'type_content.tmpl', mode='r').read()
        schema = api_content.get('types')
        # SUBCLASS_DICT = {}
        for name, item in schema.items():
            subclasses = item.get('subtypes')
            comments = "// " + "\n// ".join(item.get('description'))
            fields = item.get('fields')
            if subclasses and len(subclasses) != 0:
                field_text = get_inheritance(subclasses, SUBCLASS_DICT, schema, name)
                content += content_temp.format(
                    name=name,
                    mode="struct",
                    comments=comments,
                    fields=field_text[:-1]
                )

            elif fields is None:
                content += content_temp.format(
                    name=name,
                    mode="interface",
                    comments=comments,
                    fields=""
                )
            else:
                field_text = ''
                for field in fields:
                    text = ''
                    for types in field.get('types'):
                        field_name = field.get('name')
                        def_types = get_type(types)
                        if "InputFile" in def_types:
                            continue
                        if field_name == "chat_id" and def_types == "string":
                            continue
                        text += get_field_text(field_name, def_types, field)
                    field_text += text

                content += content_temp.format(
                    name=name,
                    mode="struct",
                    comments=comments,
                    fields=field_text[:-1]
                )

                if SUBCLASS_DICT.get(name):
                    method = SUBCLASS_DICT.get(name)
                    content += subclass_temp.format(
                        class_name=name,
                        method=method
                    )

                if name in MARKUP:
                    content += markup_temp.format(classname=name)
        content += markup_content
    return content


def build_methods(api_content: Union[dict, None]):
    content = ''
    if api_content is not None:
        schema = api_content.get('methods')
        content_temp = open(TEMPLATE / 'methods_content.tmpl', mode='r').read()
        for name, item in schema.items():
            comments = "// " + "\n// ".join(item.get('description'))
            fields = item.get('fields')
            t_returns, ex = get_field_type(item.get('returns')[0])
            if t_returns == "bool":
                is_bool = 'false'
            elif t_returns == "string":
                is_bool = "\"\""
            elif t_returns == "int64":
                is_bool = "0"
            else:
                is_bool = 'nil'
            returns = ex + t_returns + ", "
            single_returns = construct_returns(item.get('returns')[0], t_returns)
            if fields:
                field_params, method_struct_list, required, optional = method_fields(fields, name)
                struct_params = construct_opts(method_struct_list)
                struct_comment = f'// {camel(name)} methods\'s optional params'
                params = construct_params(required, optional, is_bool)

                if len(method_struct_list) > 0:
                    content += content_temp.format(
                        struct_comment=struct_comment,
                        struct_name=camel(name),
                        method_name=name,
                        struct_params=struct_params,
                        comments=comments,
                        fields=field_params,
                        returns=returns,
                        single_returns=single_returns,
                        params=params,
                        is_bool=is_bool
                    )
                else:
                    content += nofield_method_temp.format(
                        comments=comments,
                        struct_name=camel(name),
                        method_name=name,
                        fields=field_params,
                        returns=returns,
                        single_returns=single_returns,
                        params=params,
                        is_bool=is_bool
                    )
            else:
                content += nofield_method_temp.format(
                    comments=comments,
                    struct_name=camel(name),
                    method_name=name,
                    fields="",
                    returns=returns,
                    single_returns=single_returns,
                    params="",
                    is_bool=is_bool
                )

    return content


if __name__ == '__main__':
    api = generate_api()
    type_content = build_types(api)
    type_content += get_unmarshals()
    write_types(type_content)
    method_content = build_methods(api)
    write_methods(method_content)

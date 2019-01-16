namespace go example
namespace php example
namespace cpp example
namespace java example

struct Data {
    1: string text
}

service format_data {
    Data do_format(1:Data data),
}


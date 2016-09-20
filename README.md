### Go support functions to make live easier

#### Packages and its functions


* OS support ```github.com/aboltart/go-support/os```

        ExistsPath - Check if exists specified path. Returns true or false.
        IsDirectory - Check if directory exists. Returns true or false.
        IsFile - Check if file exists. Returns true or false.

* String support ```github.com/aboltart/go-support/string```

        CharOfString - Get specified by position character from string. Returns string.
        SnakeToCamel - Convert snake case string to camel case. Returns string.
        CamelToSnake - Convert camel case string to snake case. Returns string.

* Slice support ```github.com/aboltart/go-support/slice```

        StringInSlice - Check if passed string exists into slice. Returns true or false.
        CompactStringSlice - From slice remove empty string values. Returns slice.
        RemoveFromSlice - Remove from slice specified string value. Returns slice.
        PrependForSlice - Prepend string value in existing slice. Returns slice.

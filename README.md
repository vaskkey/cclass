# CCLASS

Quick util script to generate `*.hpp` and `*.cpp` files with classes in namespaces

### Usage

```
cclass namespace/another_namespace/.../class_name
```

Will generate files in `namespace/another_namespace/.../class_name` folder with 

Hpp:

```cpp
#pragma once

namespace namespace {
namespace another_namespace {

class class_name{
	public:
	class_name();
};

}
}

```

Cpp:

```cpp
#include "class_name.hpp"


namespace namespace {
namespace another_namespace {

class_name::class_name ()
{
}

}
}

```

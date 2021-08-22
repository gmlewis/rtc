# rtc

rtc is an experimental fun project of working through [The Ray Tracer
Challenge](https://pragprog.com/titles/jbtracer/the-ray-tracer-challenge/)
book by Jamis Buck, implemented in Go.

## Status

This project is still experimental but completely implements the book's
ray tracer in Go.

## Usage

```bash
go run cmd/test-obj/main.go -xsize 1280 -ysize 1024 file.obj
go run cmd/test-yaml/main.go -xsize 1280 -ysize 1024 file.yaml
```

## Examples

Here is the cover scene of the book as described in [examples/cover/cover.yaml]:

![Cover](examples/cover/cover.png)

----------------------------------------------------------------------

**Enjoy!**

----------------------------------------------------------------------

# License

Copyright 2021 Glenn M. Lewis. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

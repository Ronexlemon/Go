module values

go 1.21.3

replace greetings => ../../goModule/greetings

require (
	example/pointers v0.0.0-00010101000000-000000000000
	multimodule v0.0.0-00010101000000-000000000000
)

require golang.org/x/example/hello v0.0.0-20231013143937-1d6d2400d402 // indirect

replace error/handle => ../../ErrorHandle

replace randomGreeting => ../../random

replace multimodule => ../../multi-module/workspace/hello

replace example/pointers => ../../pointers

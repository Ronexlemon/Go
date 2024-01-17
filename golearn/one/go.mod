module values

go 1.21.3

replace greetings => ../../goModule/greetings

require (
	error/handle v0.0.0-00010101000000-000000000000
	greetings v0.0.0-00010101000000-000000000000
)

replace error/handle => ../../ErrorHandle

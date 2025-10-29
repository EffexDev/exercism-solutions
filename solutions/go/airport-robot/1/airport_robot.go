package airportrobot

type Greeter interface {
    LanguageName() string
    Greet(visitorName string) string
}

type German struct{}

func (g German) LanguageName() string {
	return "German"
}

func (g German) Greet(visitorName string) string {
    return "Hallo " + visitorName
}

type Italian struct{}

func (i Italian) LanguageName() string {
    return "Italian:"
}

func (i Italian) Greet(visitorName string) string {
    return "I can speak Italian: Ciao " + visitorName + "!"
}

type Portuguese struct{}

func (p Portuguese) LanguageName() string {
    return "Portuguese"
}

func (p Portuguese) Greet(visitorName string) string {
    return "I can speak Portuguese: Olá " + visitorName + "!"
}

func SayHello(visitorName string, greeter Greeter) string {
	return greeter.Greet(visitorName)
}
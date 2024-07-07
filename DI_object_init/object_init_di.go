package main

// ServiceImpl는 OrderService 인터페이스를 구현하는 구체적인 타입
type ServiceImpl struct{}

// Apply 메서드는 단순히 nil을 반환한다.
func (s *ServiceImpl) Apply(id int) error { return nil }

// OrderService 인터페이스는 Apply 메서드를 정의한다.
// 이는 이 인터페이스를 구현하는 모든 타입이 이 메서드를 제공해야 함을 의미한다.
type OrderService interface {
	Apply(int) error
}

// Application은 OrderService 타입의 필드를 가지고 있다.
// 이를 통해 Application은 OrderService 인터페이스를 통해 메서드를 호출할 수 있다.
type Application struct {
	os OrderService
}

// NewApplication 함수는 OrderService 인터페이스를 인자로 받아 Application 타입의 인스턴스를 생성한다.
// 이는 다른 언어에서의 생성자 주입(Constructor Injection)에 해당한다.
func NewApplication(os OrderService) *Application {
	return &Application{os: os}
}

// Application 타입의 Apply 메서드는 내부의 OrderService 인터페이스의 Apply 메서드를 호출한다.
func (app *Application) Apply(id int) error {
	return app.os.Apply(id)
}

// main 함수에서는 ServiceImpl의 인스턴스를 생성하고 이를 NewApplication 함수에 전달하여 Application 인스턴스를 생성한다.
// 그리고 Apply 메서드를 호출한다.
func main() {
	app := NewApplication(&ServiceImpl{})
	app.Apply(19)
}

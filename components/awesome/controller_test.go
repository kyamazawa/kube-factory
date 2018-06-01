package awesome

import "testing"

func TestNewController(t *testing.T) {
	t.Run("Controller has default interactor", func(t *testing.T) {
		// Act
		controller := NewController()

		// Assert
		actual := controller.interactor
		if actual == nil {
			t.Errorf("got: %v\nwant: %v", actual, "NotNil")
		}

	})

	t.Run("The interactor is Interaction interface compatible", func(t *testing.T) {
		// Arrange
		expected := true

		// Act
		controller := NewController()

		// Assert
		_, actual := controller.interactor.(Interaction)
		if actual != expected {
			t.Errorf("got: %v\nwant: %v", actual, expected)
		}
	})

	t.Run("The interactor is injectable", func(t *testing.T) {
		// Arrange
		i1 := NewInteractor()
		i2 := NewInteractor()

		// Act
		controller := NewController(WithInteractor(i1))

		// Assert
		actual := controller.interactor
		if actual != i1 {
			t.Errorf("got: %v\nwant: %v", actual, i1)
		}
		if actual == i2 {
			t.Errorf("got: %v\nwant: %v", actual, i2)
		}
	})

	t.Run("The interactor should not be nil with wrong usage", func(t *testing.T) {
		// Arrange
		var i1 *Interactor // nil

		// Act
		controller := NewController(WithInteractor(i1))

		// Assert
		actual := controller.interactor
		if actual == nil {
			t.Errorf("got: %v\nwant: %v", actual, "NotNil")
		}
	})
}

func TestSomeUsecase(t *testing.T) {
	t.Run("interactor.DoSomeUsecase should be called", func(t *testing.T) {
		// Arrange
		spy := &TestSomeUsecaseSpy{DoSomeUsecaseCalled: false}
		controller := NewController(WithInteractor(spy))

		// Act
		controller.SomeUsecase()

		// Assert
		actual := spy.DoSomeUsecaseCalled
		if actual != true {
			t.Errorf("got: %v\nwant: %v", actual, "DoSomeUsecase does not called")
		}
	})

	t.Run("it should response SomeUsecaseViewModel", func(t *testing.T) {
		// Arrange
		mockVM := &SomeUsecaseViewModel{}
		spy := &TestSomeUsecaseSpy{MockViewModel: mockVM}
		controller := NewController(WithInteractor(spy))

		// Act
		actual, err := controller.SomeUsecase()

		// Assert
		if actual != mockVM {
			t.Errorf("got: %v\nwant: %v", actual, mockVM)
		}
		if err != nil {
			t.Errorf("got: %v\nwant: %v", err, "Should be nil")
		}
	})
}

type TestSomeUsecaseSpy struct {
	DoSomeUsecaseCalled bool
	MockViewModel       *SomeUsecaseViewModel
}

func (s *TestSomeUsecaseSpy) DoSomeUsecase(req *SomeUsecaseRequest, callback func(*SomeUsecaseViewModel, error)) {
	s.DoSomeUsecaseCalled = true
	callback(s.MockViewModel, nil)
}

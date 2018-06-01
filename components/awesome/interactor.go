package awesome

// Interactor is ...
type Interactor struct {
	presenter Presentation
}

// InteractorOption is ...
type InteractorOption func(*Interactor)

// WithPresenter is ...
func WithPresenter(p *Presenter) InteractorOption {
	return func(s *Interactor) {
		if p != nil {
			s.presenter = p
		}
	}
}

// NewInteractor is ...
func NewInteractor(opts ...InteractorOption) *Interactor {
	interactor := &Interactor{}

	for _, opt := range opts {
		opt(interactor)
	}

	if interactor.presenter == nil {
		presenter := NewPresenter()
		interactor.presenter = presenter
	}

	return interactor
}

// Interaction is ...
type Interaction interface {
	DoSomeUsecase(req *SomeUsecaseRequest, callback func(*SomeUsecaseViewModel, error))
}

// DoSomeUsecase is ...
func (s *Interactor) DoSomeUsecase(req *SomeUsecaseRequest, callback func(*SomeUsecaseViewModel, error)) {
	res := &SomeUsecaseResponse{}
	s.presenter.PresentSomeUsecase(res, callback)
}

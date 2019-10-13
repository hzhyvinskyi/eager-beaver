package dal

type Dal interface {
    User() UserRepository
}

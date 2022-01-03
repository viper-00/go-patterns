package iterator

// wiki: https://en.wikipedia.org/wiki/Iterator_pattern
//
// Provide a way to access the elements of an aggregate object sequentially without exposing its underlying representation.

type user struct {
	name string
	age  int
}

type Collection interface {
	createIterator() iterator
}

type userCollection struct {
	users []*user
}

func (u *userCollection) createIterator() iterator {
	return &userIterator{
		users: u.users,
	}
}

type iterator interface {
	hasNext() bool
	getNext() *user
}

type userIterator struct {
	index int
	users []*user
}

func (u *userIterator) hasNext() bool {
	return u.index < len(u.users)
}

func (u *userIterator) getNext() *user {
	if u.hasNext() {
		user := u.users[u.index]
		u.index++
		return user
	}
	return nil
}

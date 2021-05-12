package iterator

// wiki: https://en.wikipedia.org/wiki/Iterator_pattern

/**
 * Iterator is a behavioral design pattern that traverse elements of a collection
 * without exposing its underlying representation(list, stack, tree, etc).
 */

// client code - user
type user struct {
	name string
	age  int
}

// collection
type collection interface {
	createIterator() iterator
}

// concrete collection - userCollection
type userCollection struct {
	users []*user
}

func (u *userCollection) createIterator() iterator {
	return &userIterator{
		users: u.users,
	}
}

// iterator
type iterator interface {
	hasNext() bool
	getNext() *user
}

// concrete iterator - userIterator
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

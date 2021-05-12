package iterator

import "testing"

func TestPattern(t *testing.T) {
	u1 := &user{name: "zmy", age: 21}
	u2 := &user{name: "abc", age: 30}

	userCollection := &userCollection{users: []*user{u1, u2}}

	iterator := userCollection.createIterator()

	for iterator.hasNext() {
		user := iterator.getNext()
		t.Logf("current user name is %s and age is %d\n", user.name, user.age)
	}
}

package fixtures

const (
	QUERY_NAME = "HeroNameAndFriends"
)

var Query string = `
query HeroNameAndFriends {
	hero {
	  name
	  friends {
		name
	  }
	}
}
`

var QueryWithNoName string = `
{
	hero {
	  name
	  friends {
		name
	  }
	}
}
`

var InvalidQuery string = `
query HeroNameAndFriends {
	hero 
	  name
	  friends {
		name
	  }
	}
}
`

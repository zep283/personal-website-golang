package github

// curl -L -X POST 'https://api.github.com/graphql' \
// -H 'Authorization: bearer <token>' \
// --data-raw '{"query":"{\n  user(login: \"GabrielBB\") {\n pinnedItems(first: 6, types: REPOSITORY) {\n nodes {\n ... on Repository {\n name\n }\n }\n }\n }\n}"
// '

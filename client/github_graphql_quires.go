package client

const githubListOpenPullRequestsQuery = `
query ListOpenPullRequests($owner: String!, $repo: String!) {
  repository(owner: $owner, name: $repo) {
    pullRequests(states: OPEN, first: 100) {
      nodes {
        number
        title
        url
        mergeable
        reviewDecision
        isInMergeQueue
        author {
          login
        }
        headRefName
        commits(last: 1) {
          nodes {
            commit {
              statusCheckRollup {
                state
              }
            }
          }
        }
      }
    }
  }
}`

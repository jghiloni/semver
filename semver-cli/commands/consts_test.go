package commands_test

const fakeStdIn = `0.0.4
1.2.3 10.20.30 1.1.2-prerelease+meta


1.1.2+meta
1.0.0-alpha
1.0.0-beta
1.0.0-alpha.beta
1.0.0-alpha.beta.1
1.0.0-alpha.1
1.0.0-alpha0.valid
1.0.0-alpha.0valid
1.0.0-alpha-a.b-c-somethinglong+build.1-aef.1-its-okay
1.0.0-rc.1+build.1  2.0.0-rc.1+build.123
1.2.3-beta
10.2.3-DEV-SNAPSHOT
1.2.3-SNAPSHOT-123`

const normalized = `0.0.4
1.2.3
10.20.30
1.1.2-prerelease+meta
1.1.2+meta
1.0.0-alpha
1.0.0-beta
1.0.0-alpha.beta
1.0.0-alpha.beta.1
1.0.0-alpha.1
1.0.0-alpha0.valid
1.0.0-alpha.0valid
1.0.0-alpha-a.b-c-somethinglong+build.1-aef.1-its-okay
1.0.0-rc.1+build.1
2.0.0-rc.1+build.123
1.2.3-beta
10.2.3-DEV-SNAPSHOT
1.2.3-SNAPSHOT-123`

const sortedAsc = `0.0.4
1.0.0-alpha
1.0.0-alpha.1
1.0.0-alpha.0valid
1.0.0-alpha.beta
1.0.0-alpha.beta.1
1.0.0-alpha-a.b-c-somethinglong+build.1-aef.1-its-okay
1.0.0-alpha0.valid
1.0.0-beta
1.0.0-rc.1+build.1
1.1.2-prerelease+meta
1.1.2+meta
1.2.3-beta
1.2.3-SNAPSHOT-123
1.2.3
2.0.0-rc.1+build.123
10.2.3-DEV-SNAPSHOT
10.20.30`

const sortedDesc = `10.20.30
10.2.3-DEV-SNAPSHOT
2.0.0-rc.1+build.123
1.2.3
1.2.3-SNAPSHOT-123
1.2.3-beta
1.1.2+meta
1.1.2-prerelease+meta
1.0.0-rc.1+build.1
1.0.0-beta
1.0.0-alpha0.valid
1.0.0-alpha-a.b-c-somethinglong+build.1-aef.1-its-okay
1.0.0-alpha.beta.1
1.0.0-alpha.beta
1.0.0-alpha.0valid
1.0.0-alpha.1
1.0.0-alpha
0.0.4`

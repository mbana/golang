# `READ-WRITE`

Please look at Confluence page at <https://openbanking.atlassian.net/wiki/spaces/DZ/pages/23363964/Read+Write+Data+API+Specifications>.

Download links:

* <https://github.com/OpenBankingUK/read-write-api-specs>.
* <https://github.com/OpenBankingUK/account-info-api-spec>.

## Copy Specifications

**TODO:** Tidyup script to get all the specifications.

```sh
git clone git@github.com:OpenBankingUK/read-write-api-specs.git
cd read-write-api-specs
(VER="v3.1.1"; git checkout master && git checkout $VER && git --no-pager branch && cp -rv ./dist/* /Users/mbana/work/openbankinguk/bitbucket/openbanking_tools/specifications/read-write/${VER}/)
(VER="v3.1.2"; git checkout master && git checkout $VER && git --no-pager branch && cp -rv ./dist/* /Users/mbana/work/openbankinguk/bitbucket/openbanking_tools/specifications/read-write/${VER}/)
(VER="v3.1.3"; git checkout master && git checkout $VER && git --no-pager branch && cp -rv ./dist/* /Users/mbana/work/openbankinguk/bitbucket/openbanking_tools/specifications/read-write/${VER}/)
(VER="v4.0"; git checkout master && git checkout $VER && git --no-pager branch && cp -rv ./dist/* /Users/mbana/work/openbankinguk/bitbucket/openbanking_tools/specifications/read-write/${VER}/)
```

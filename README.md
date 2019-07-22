# SRP module for CMaNGOS

this module implements the SRP authentication mechanism for CMaNGOS cores.
It can be integrated into go-based content management systems.

## example client

the example client creates `s` and `v` values that can be insert into account table of the
realmd database.

    > ./wowpasswd johndoe
    Password: 
    Retype Password: 
    s: 34508ECC8253DEEF06EB1FA4FEE7645BED7E77BE26D3D41819BEDB9C6B71F619
    v: 6D699DCB5C39F693FB301BB633FFCEEF55884CBD87C1E52FC768890CFA769311

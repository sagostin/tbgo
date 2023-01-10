## Information

This is a tool designed to be used with the TelcoBridges Free/ProSBC and to tie into their REST API. It's mainly used to manage the NAP creation process along with adding numbers.

## Features

* [X]  Creation of NAPs (some modifications to the code may be required to match your exact setup)
* [X]  Creation / Modification of files in the File DB
* [X]  Command Line Arguments
* [X]  Importable Go package
* [X]  NAP Column values
* [ ]  Good CLI usage with multiple options/features to manage/edit the SBC

## Installation & Usage

Requirements: Go 1.19

1. Clone this repository
2. Build the application using `go build` while in the directory
3. Usage:

   ```
   tbgo.exe --host https://host:port --username USERNAME --password PASSWORD --napcreate --pbx --customer=WadesWindowWashing --napproxyhost=192.168.0.1:5060 --numbers=5555555555,5555555522 --config=config_1 --portrange=Host.pr_LAN0 --siptransport=LAN0_5060 --digitmap=digitmap.csv --napcroutegroups=55,11,12,32 --rdefroutegroups=55 --napprofile=Zultys
   ```

## Arguments / Flags

* `--host https://0.0.0.0:12358` *telcobridges api address*
* `--username root` *api username*
* `--password P@ssw0rd` *api password*
* `--napcreate` *defines wether to create a nap based on the provided flags*
* `--pbx` *defines if a nap is a pbx nap*
* `--customer BobsBurgers` *name to use when creating nap files and such*
* `--napproxyhost 192.168.0.1:5060` *endpoint for nap*
* `--numbers 5555555557,5555555558` *phone numbers to use in nap creation*
* `--config config_1` *config to make changes to*
* `--portrange Host.pr_WAN0` *define the portrange to use when creating a nap, it will likely always start with* **Host.PORT_RANGE_NAME**
* `--siptransport WAN0_5060` *name of the sip transport to be used*
* `--digitmap digitmap.csv` *digitmap to be used*
* `--rdefroutegroups 55,10,1` *route groups to be used in routedef creation*
* `--napcroutegroups 55,10,1` *route groups to be used in nap colum creation*
* `--napprofile default` *specify the nap profile*

## Licence

[`GNU Affero General Public License v3.0`](https://github.com/sagostin/netwatcher-agent/blob/master/LICENSE.md)

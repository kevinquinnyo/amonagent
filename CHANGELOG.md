0.4.9 - 03.02.2016
==============

* Telegraf Plugin

0.4.8 - 22.01.2016
==============

* Fixes an issue where the agent will hang and not send data on specific servers

0.4.7 - 21.01.2016
==============

* CloudID fixes


0.4.6 - 20.01.2016
==============

* Postinst script fixes
* Improved test command
* PostgreSQL plugin - slow queries parser fix

0.4.5 - 19.01.2016
==============

* Missing Full Name for database/index size in the MySQL Plugin(Triggers an error in Amon)


0.4.4 - 15.01.2016
==============

* Recompile with Go 1.5.3(Fixes security-related issue in Go) - https://groups.google.com/forum/#!topic/golang-dev/MEATuOi_ei4

0.4.3 - 13.01.2016
==============

* Properly format process memory metric
* Properly format disc metrics for values bigger than Terrabyte

0.4.2 - 09.01.2016
==============

* Generate machine id on first install

0.4.1 - 08.01.2016
==============

* Fix init script on systemd distros

0.4 - 07.01.2016
==============

* Collects all metrics in parallel
* MySQL Plugin
* PostgreSQL Plugin
* MongoDB Plugin
* Redis Plugin
* HAProxy Plugin
* Sensu Plugin
* Nginx Plugin
* Apache Plugin
* Health Checks Plugin
* Can run health checks locally and send the results to Amon
* Custom Plugin - you can write custom plugins in any language with just a couple lines of code.
* New command line options - `list-plugins`, `test-plugin`, `plugin-config`
* Gets Amazon, Google and DigitalOcean instance ids
* Works with self-signed certificates(skips TSL verification)

0.3 - 20.12.2015
==============

* More detailed error messages
* Improve testing command

0.2.5 - 17.12.2015
==============

* Fix permissions issues in the systemd service file

0.2.1 - 15.12.2015
==============

* Machine id parameter
* Fix CPU collector, format data to float

0.2 - 14.12.2015
==============

* Initial release

Solr monitoring agent for NewRelic
===============
[![Build Status](https://travis-ci.org/yvasiyarov/newrelic_solr.png)](https://travis-ci.org/yvasiyarov/newrelic_solr)

Installation
-------------

If you have not Go compiler in your system:   
`sudo apt-get install golang`  

Install dependencies:   
`sudo go get github.com/yvasiyarov/newrelic_platform_go`   

Get and build agent:   
`git clone https://github.com/yvasiyarov/newrelic_solr.git   
cd newrelic_solr   
go build -o solr_agent`   

Run agent in debug mode:   
`./solr_agent --verbose=true --solr-url="127.0.0.1:8080/" --newrelic-license=[your newrelic license key]`   

In production mode you can run it with nohup:  
`nohup ./sphinx_agent --solr-url="127.0.0.1:8080/" --newrelic-license=[your newrelic license key]`  



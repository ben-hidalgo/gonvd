FROM golang:1.12.1-stretch

ENV GO111MODULE=on

WORKDIR /usr/local/cve

RUN wget https://nvd.nist.gov/feeds/json/cve/1.0/nvdcve-1.0-2002.json.gz
RUN wget https://nvd.nist.gov/feeds/json/cve/1.0/nvdcve-1.0-2003.json.gz
RUN wget https://nvd.nist.gov/feeds/json/cve/1.0/nvdcve-1.0-2004.json.gz
RUN wget https://nvd.nist.gov/feeds/json/cve/1.0/nvdcve-1.0-2005.json.gz
RUN wget https://nvd.nist.gov/feeds/json/cve/1.0/nvdcve-1.0-2006.json.gz
RUN wget https://nvd.nist.gov/feeds/json/cve/1.0/nvdcve-1.0-2007.json.gz
RUN wget https://nvd.nist.gov/feeds/json/cve/1.0/nvdcve-1.0-2008.json.gz
RUN wget https://nvd.nist.gov/feeds/json/cve/1.0/nvdcve-1.0-2009.json.gz
RUN wget https://nvd.nist.gov/feeds/json/cve/1.0/nvdcve-1.0-2010.json.gz
RUN wget https://nvd.nist.gov/feeds/json/cve/1.0/nvdcve-1.0-2011.json.gz
RUN wget https://nvd.nist.gov/feeds/json/cve/1.0/nvdcve-1.0-2012.json.gz
RUN wget https://nvd.nist.gov/feeds/json/cve/1.0/nvdcve-1.0-2013.json.gz
RUN wget https://nvd.nist.gov/feeds/json/cve/1.0/nvdcve-1.0-2014.json.gz
RUN wget https://nvd.nist.gov/feeds/json/cve/1.0/nvdcve-1.0-2015.json.gz
RUN wget https://nvd.nist.gov/feeds/json/cve/1.0/nvdcve-1.0-2016.json.gz
RUN wget https://nvd.nist.gov/feeds/json/cve/1.0/nvdcve-1.0-2017.json.gz
RUN wget https://nvd.nist.gov/feeds/json/cve/1.0/nvdcve-1.0-2018.json.gz
RUN wget https://nvd.nist.gov/feeds/json/cve/1.0/nvdcve-1.0-2019.json.gz

RUN wget https://nvd.nist.gov/feeds/json/cve/1.0/nvdcve-1.0-modified.json.gz
RUN wget https://nvd.nist.gov/feeds/json/cve/1.0/nvdcve-1.0-recent.json.gz

RUN gunzip *.gz

#
WORKDIR /usr/local/gocache

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download


WORKDIR /usr/local/gonvd

COPY go.mod           go.mod
COPY go.sum           go.sum
COPY entrypoint.sh    entrypoint.sh
COPY main.go          main.go
COPY app/             app/
COPY restful/         restful/

EXPOSE 8000

CMD [ "bash", "entrypoint.sh" ]

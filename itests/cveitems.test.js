import test from 'ava'
import http from 'ava-http'

const hostname=process.env.HOSTNAME
const port=    process.env.PORT

// The "recent" file contains duplicated IDs which overwrite the modified entries
// $ cat ignored/nvdcve-1.0-modified.json | jq '.CVE_Items' | grep CVE_data_meta | wc -l
//     1620
// $ cat ignored/nvdcve-1.0-recent.json | jq '.CVE_Items' | grep CVE_data_meta | wc -l
//      489


test('GET /cveitems', async t => {
    await http.get(`http://${hostname}:${port}/cveitems`)
        .then(res => {
            t.is(1620, res.length)
        })
})

test('GET /cveitems/{id} missing id', async t => {
    await http.get(`http://${hostname}:${port}/cveitems/notanid`)
        .catch(err => {
            t.is(err.statusCode, 404)
        })
})

test('GET /cveitems/{id}', async t => {
    await http.get(`http://${hostname}:${port}/cveitems/CVE-2013-2805`)
        .then(res => {
            t.is(res.id, 'CVE-2013-2805')
        })
})

import test from 'ava'
import http from 'ava-http'

const hostname=process.env.HOSTNAME
const port=    process.env.PORT

test('GET /cveitems', async t => {
    await http.get(`http://${hostname}:${port}/cveitems`)
        .then(res => {
            //TODO: change the length / test to handle pagination
            t.is(2109, res.length)
        })
})

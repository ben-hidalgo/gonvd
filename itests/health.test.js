import test from 'ava'
import http from 'ava-http'

const hostname=process.env.HOSTNAME
const port=    process.env.PORT

test('GET /', async t => {
    await http.get(`http://${hostname}:${port}/`)
        .then(res => {
            t.deepEqual(res, {})
        })
})

test('GET /health', async t => {
    await http.get(`http://${hostname}:${port}/health`)
        .then(res => {
            t.deepEqual(res, {status: 'UP'})
        })
})

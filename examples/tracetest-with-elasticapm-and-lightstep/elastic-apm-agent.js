// Add this to the VERY top of the first file loaded in your app
const apm = require('elastic-apm-node').start({
  serviceName: 'sample-app',
  serverUrl: 'http://apm-server:8200',
})

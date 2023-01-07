var Base = require('../base');
var https = require('https');
var querystring = require('querystring');


class Rest extends Base {
    
    constructor(username, password) {
        super(username, password);
    }

    request(method, params) {
        let data = Object.assign({}, this.data, params);
        var postdata = querystring.stringify(data);

        var options = {
            host: 'rest.payamak-panel.com',
            port: 443,
            path: `/api/SendSMS/${method}`,
            method: 'POST',
            headers: {
                'Content-Length': postdata.length,
                'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8',
            }
        };

        var req = https.request(options, function(res) {
            console.log('STATUS: ' + res.statusCode);
            res.setEncoding('utf8');
            res.on('data', function(chunk) {
                console.log(chunk);
            });
        });

        req.on('error', function(e) {
            console.log('problem with request: ' + e.message);
        });

        req.write(postdata, "utf8");
        req.end();
    }


    send(to, from, text, isFlash = false) {
        return this.request('SendSMS', {
            to,
            from: from,
            text,
            isFlash
        });
    }

    sendByBaseNumber(text, to, bodyId) {
        return this.request('BaseServiceNumber', {
            text,
            to,
            bodyId
        });
    }

    isDelivered(recId) {
        return this.request('GetDeliveries2', {
            recId
        });
    }

    getMessages(location, index, count, from = '') {
        return this.request('GetMessages', {
            location,
            index,
            count
        });
    }

    getCredit() {
        return this.request('GetCredit', {});
    }

    getBasePrice() {
        return this.request('GetBasePrice', {});
    }
    
    getNumbers() {
        return this.request('GetUserNumbers', {});
    }

}

module.exports = Rest;
/**
* IPService.js
*
* @description :: Service to retrieve the public IP address for the node.
*/

var http = require('http');

module.exports = {

    publicIPv4Address: function(cb) {
        try {
            var req = http.request({ hostname: 'ipv4.icanhazip.com' }, function(res) {
                res.on('data', function (body) {
                    var result = new String(body).trim();
                    cb(null, result);
                });
            });
            req.on('error', function(e) {
                throw new Error(e.message);
            });
            req.end();
        } catch(err) {
            cb(err, null);
        }
    },

    publicIPv6Address: function(cb) {
        try {
            var req = http.request({ hostname: 'ipv6.icanhazip.com' }, function(res) {
                res.on('data', function (body) {
                    var result = new String(body).trim();
                    cb(null, result);
                });
            });
            req.on('error', function(e) {
                throw new Error(e.message);
            });
            req.end();
        } catch(err) {
            cb(err, null);
        }
    },

    updatePublicAddress: function() {
        this.publicIPv4Address(function(err, ip) {
            if (err) {
                console.error(err)
                return;
            }

            Address.findOne({ipv4: ip}).exec(function (err, found) {
                if (found) {
                    found.save(function (err, address) {
                        if (err) {
                            console.error('Unable to update Address: ' + err);
                        } else {
                            console.log('Updated IPv4 Address: ' + address.ipv4);
                        }
                    });
                } else {
                    Address.create({ipv4: ip}).exec(function(err, created) {
                        if (err) {
                            console.error('Unable to create Address: ' + err);
                        } else {
                            console.log('Created IPv4 address: ' + created.ipv4);
                            EmailService.sendIPUpdateEmail(created);
                        }
                    });
                }
            });
        });
    }

};

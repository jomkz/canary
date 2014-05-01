var http = require('http');

module.exports = {
	please: function (cb) {
		try {
			var req = http.request({ hostname: 'ipv4.icanhazip.com' }, function(res) {
				res.on('data', function (body) {
					cb(null, body);
				});
			});
			req.on('error', function(e) {
				throw new Error(e.message);
			});
			req.end();
		} catch(err) {
			cb(err, null);
		}
	}
};

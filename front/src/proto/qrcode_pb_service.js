// package: qrcode
// file: qrcode.proto

var qrcode_pb = require("./qrcode_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var QRCodeService = (function () {
  function QRCodeService() {}
  QRCodeService.serviceName = "qrcode.QRCodeService";
  return QRCodeService;
}());

QRCodeService.GetURL = {
  methodName: "GetURL",
  service: QRCodeService,
  requestStream: false,
  responseStream: false,
  requestType: qrcode_pb.Empty,
  responseType: qrcode_pb.URL
};

QRCodeService.UpdateURL = {
  methodName: "UpdateURL",
  service: QRCodeService,
  requestStream: false,
  responseStream: false,
  requestType: qrcode_pb.URL,
  responseType: qrcode_pb.Empty
};

QRCodeService.GetQRCodes = {
  methodName: "GetQRCodes",
  service: QRCodeService,
  requestStream: false,
  responseStream: false,
  requestType: qrcode_pb.Empty,
  responseType: qrcode_pb.QRCodeList
};

QRCodeService.AddQRCodes = {
  methodName: "AddQRCodes",
  service: QRCodeService,
  requestStream: false,
  responseStream: false,
  requestType: qrcode_pb.Empty,
  responseType: qrcode_pb.Empty
};

exports.QRCodeService = QRCodeService;

function QRCodeServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

QRCodeServiceClient.prototype.getURL = function getURL(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(QRCodeService.GetURL, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

QRCodeServiceClient.prototype.updateURL = function updateURL(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(QRCodeService.UpdateURL, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

QRCodeServiceClient.prototype.getQRCodes = function getQRCodes(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(QRCodeService.GetQRCodes, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

QRCodeServiceClient.prototype.addQRCodes = function addQRCodes(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(QRCodeService.AddQRCodes, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

exports.QRCodeServiceClient = QRCodeServiceClient;


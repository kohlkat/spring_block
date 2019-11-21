"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.getLinkPosition = function (node, portId) {
    var port = node.ports[portId];
    return {
        x: node.position.x + (port.position ? port.position.x : 0),
        y: node.position.y + (port.position ? port.position.y : 0),
    };
};
//# sourceMappingURL=getLinkPosition.js.map
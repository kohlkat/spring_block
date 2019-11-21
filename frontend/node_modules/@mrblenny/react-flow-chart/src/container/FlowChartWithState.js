"use strict";
var __extends = (this && this.__extends) || (function () {
    var extendStatics = function (d, b) {
        extendStatics = Object.setPrototypeOf ||
            ({ __proto__: [] } instanceof Array && function (d, b) { d.__proto__ = b; }) ||
            function (d, b) { for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p]; };
        return extendStatics(d, b);
    };
    return function (d, b) {
        extendStatics(d, b);
        function __() { this.constructor = d; }
        d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
    };
})();
Object.defineProperty(exports, "__esModule", { value: true });
var React = require("react");
var __1 = require("../");
var actions = require("./actions");
var mapValues_1 = require("./utils/mapValues");
/**
 * Flow Chart With State
 */
var FlowChartWithState = /** @class */ (function (_super) {
    __extends(FlowChartWithState, _super);
    function FlowChartWithState(props) {
        var _this = _super.call(this, props) || this;
        _this.stateActions = mapValues_1.default(actions, function (func) {
            return function () {
                var args = [];
                for (var _i = 0; _i < arguments.length; _i++) {
                    args[_i] = arguments[_i];
                }
                return _this.setState(func.apply(void 0, args));
            };
        });
        _this.state = props.initialValue;
        return _this;
    }
    FlowChartWithState.prototype.render = function () {
        var _a = this.props, Components = _a.Components, config = _a.config;
        return (React.createElement(__1.FlowChart, { chart: this.state, callbacks: this.stateActions, Components: Components, config: config }));
    };
    return FlowChartWithState;
}(React.Component));
exports.FlowChartWithState = FlowChartWithState;
//# sourceMappingURL=FlowChartWithState.js.map
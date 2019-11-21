'use strict';

Object.defineProperty(exports, '__esModule', {
  value: true
});

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { 'default': obj }; }

var _react = require('react');

var _react2 = _interopRequireDefault(_react);

var Col = _react2['default'].createClass({
  displayName: 'Col',

  propTypes: {
    grow: _react2['default'].PropTypes.bool.isRequired,
    shrink: _react2['default'].PropTypes.bool.isRequired,
    basis: _react2['default'].PropTypes.string.isRequired,
    padding: _react2['default'].PropTypes.oneOfType([_react2['default'].PropTypes.string, _react2['default'].PropTypes.number]).isRequired,
    align: _react2['default'].PropTypes.string.isRequired
  },

  getDefaultProps: function getDefaultProps() {
    return {
      grow: true,
      shrink: true,
      basis: 'auto', // also accepts '100px', '10%', etc.
      align: 'left',
      padding: 10
    };
  },

  render: function render() {
    var styles = {
      flex: (this.props.grow ? '1 ' : '0 ') + (this.props.shrink ? '1 ' : '0 ') + this.props.basis,
      textAlign: this.props.align,
      padding: this.props.padding
    };

    return _react2['default'].createElement(
      'div',
      { style: styles },
      this.props.children
    );
  }
});
exports.Col = Col;
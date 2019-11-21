import React from 'react';

export const Row = React.createClass({
  propTypes : {
    align : React.PropTypes.string, // left or right
  },

  getDefaultProps() {
    return {
      align : 'left',
    };
  },

  render() {
    return (
      <div className={this.props.align === 'left' ? 'flexbox' : 'flexbox-right'}>
        {this.props.children}
      </div>
    );
  },
});

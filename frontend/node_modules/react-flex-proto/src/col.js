import React from 'react';

export const Col = React.createClass({
  propTypes : {
    grow : React.PropTypes.bool.isRequired,
    shrink : React.PropTypes.bool.isRequired,
    basis : React.PropTypes.string.isRequired,
    padding : React.PropTypes.oneOfType([
      React.PropTypes.string,
      React.PropTypes.number,
    ]).isRequired,
    align : React.PropTypes.string.isRequired,
  },

  getDefaultProps() {
    return {
      grow : true,
      shrink : true,
      basis : 'auto', // also accepts '100px', '10%', etc.
      align : 'left',
      padding : 10,
    };
  },

  render() {
    const styles = {
      flex : (this.props.grow ? '1 ' : '0 ') + (this.props.shrink ? '1 ' : '0 ') + this.props.basis,
      textAlign : this.props.align,
      padding : this.props.padding,
    };

    return (
      <div style={styles}>
        {this.props.children}
      </div>
    );
  },
});

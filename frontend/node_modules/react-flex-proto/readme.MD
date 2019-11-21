# React-Flex-Proto
Note: This is a partial implementation of Flex to quickly prototype responsive projects

It does not implement Flex Direction of "Col", or allow Flex items to grow by more than a scale of one.

Cells can choose to grow/shrink or not, and if not, they will remain at their initial basis size which can be automatically calculated if `basis` is not passed, otherwise you can pass in a fixed width: `basis='100px'`

## Example Usage

```javascript
import 'react-flexbox/styles/flex.css';

render() {
  return (
    <Row>
      <Col>
        <input
          type='text'
          value={this.state.value}
          onChange={this.onTextChange.bind(null, 'value')} />
      </Col>
      <Col align='right' grow={false} shrink={false}>
        <button className='btn-primary' onClick={this.onAdd}>
          <i className='fa fa-plus' /> Add
        </button>
      </Col>
    </Row>
  );
}
```
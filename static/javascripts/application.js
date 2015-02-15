var request = reqwest;

var BookCreateForm = React.createClass({
  handleSubmit: function(e) {
    e.preventDefault();

    var title       = this.refs.title.getDOMNode().value.trim();
    var description = this.refs.description.getDOMNode().value.trim();
    var color       = this.refs.color.getDOMNode().value.trim();

    request({
      url: '/api/books',
      type: 'json',
      method: 'post',
      data: {title: title, description: description, color: color},
      success: function (book) {
        // TODO notify storage
      }
    })

    this.refs.title.getDOMNode().value = '';
    this.refs.description.getDOMNode().value = '';
    this.refs.color.getDOMNode().value = '';
  },

  render: function() {
    return (
      <form onSubmit={this.handleSubmit}>
        <h1>Create book</h1>
        <p><label>Title:* <input type="text" ref="title" /></label></p>
        <p><label>Description:* <input type="text" ref="description" /></label></p>
        <p><label>Color:* <input type="text" ref="color" /></label></p>
        <input type="submit" value="Create" />
      </form>
    );
  }
});

React.render(
  <BookCreateForm />,
  document.getElementById('main')
);

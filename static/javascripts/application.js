var request = reqwest;

var BookCreateForm = React.createClass({
  handleSubmit: function(e) {
    e.preventDefault();

    var title       = this.refs.title.getDOMNode().value.trim();
    var description = this.refs.description.getDOMNode().value.trim();
    var color       = this.refs.color.getDOMNode().value.trim();


    var self = this;
    request({
      url: '/api/books',
      type: 'json',
      method: 'post',
      data: {title: title, description: description, color: color},
      success: function (book) {
        self.props.onBookSubmit(book);
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

var BookListing = React.createClass({
  render: function() {
    var bookNodes = this.props.data.map(function (book) {
      return (
        <article>
          <h1>{book.title}</h1>
          <p>{book.description}</p>
        </article>
      );
    });
    return (
      <div>
        {bookNodes}
      </div>
    );
  }
});

var BookIndex = React.createClass({
  getInitialState: function() {
    return {data: []};
  },

  componentDidMount: function() {
    request({
      url: '/api/books',
      type: 'json',
      success: function(data) {
        this.setState({data: data});
      }.bind(this),
    })
  },

  handleBookSubmit: function(book) {
    var updatedData = this.state.data.concat([book]);
    this.setState({data: updatedData});
  },

  render: function() {
    return (
      <div>
        <h1>Book Listing</h1>
        <BookListing data={this.state.data} />
        <BookCreateForm onBookSubmit={this.handleBookSubmit} />
      </div>
    );
  }
});

React.render(
  <BookIndex />,
  document.getElementById('main')
);

var request = reqwest;

var BookCreateForm = React.createClass({
  handleSubmit: function(e) {
    e.preventDefault();

    var title       = this.refs.title.getDOMNode().value.trim();
    var description = this.refs.description.getDOMNode().value.trim();
    var color       = this.refs.color.getDOMNode().value.trim();

    this.props.onBookSubmit({title: title, description: description, color: color});

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

var BookItem = React.createClass({
  handleRemoval: function() {
    if (!confirm('Are you sure?')) { return; }

    this.props.onBookRemoval(this.props.book);
  },

  render: function() {
    return (
      <article>
        <h1>{this.props.book.title}</h1>
        <p>{this.props.book.description}</p>
        <button onClick={this.handleRemoval}>delete</button>
      </article>
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
    var self = this;
    request({
      url:    '/api/books',
      type:   'json',
      method: 'post',
      data:    book,
      success: function (book) {
        var data = self.state.data.concat([book]);
        self.setState({data: data});
      }
    });
  },

  handlBookRemoval: function(book) {
    var data = this.state.data;
    _.remove(data, book);
    this.setState({data: data});
  },

  render: function() {
    var bookNodes = this.state.data.map(function (book) {
      return <BookItem book={book} onBookRemoval={this.handlBookRemoval} />;
    }, this);

    return (
      <div>
        <h1>Book Listing</h1>
        {bookNodes}
        <BookCreateForm onBookSubmit={this.handleBookSubmit} />
      </div>
    );
  }
});

React.render(
  <BookIndex />,
  document.getElementById('main')
);

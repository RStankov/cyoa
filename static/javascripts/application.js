var BookCreateForm = React.createClass({
  render: function() {
    return (
      <form action="/api/books" method="POST">
        <h1>Create book</h1>
        <p><label>Title:* <input type="text" name="title" /></label></p>
        <p><label>Description:* <input type="text" name="description" /></label></p>
        <p><label>Color:* <input type="text" name="color" /></label></p>
        <input type="submit" value="Create" />
      </form>
    );
  }
});

React.render(
  <BookCreateForm />,
  document.getElementById('main')
);

import { Link } from "react-router-dom";

function Top() {
  return (
    <div>
      <h1>Top</h1>
      <p>Welcome to the top page!</p>
      <Link to="/about">About</Link>
    </div>
  );
}
export default Top;

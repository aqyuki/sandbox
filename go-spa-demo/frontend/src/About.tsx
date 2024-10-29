import { Link } from "react-router-dom";

function About() {
  return (
    <div>
      <h1>About</h1>
      <p>
        This is a simple app to demonstrate how to use React with TypeScript.
      </p>
      <Link to="/">Top</Link>
    </div>
  );
}
export default About;

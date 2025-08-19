import Layout from "../components/Layout";
import LoginForm from "../components/LoginForm";
import Main from "../components/Main";
import NavBar from "../components/NavBar";

function Login() {
	return (
		<Layout>
			<NavBar />
			<Main>
				<LoginForm />
			</Main>
		</Layout>
	);
}

export default Login;

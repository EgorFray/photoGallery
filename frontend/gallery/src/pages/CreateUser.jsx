import CreateUserForm from "../components/CreateUserForm";
import Footer from "../components/Footer";
import Header from "../components/Header";
import Layout from "../components/Layout";
import Main from "../components/Main";
import NavBar from "../components/NavBar";

function CreateUser() {
	return (
		<Layout>
			<NavBar />
			<Main>
				<CreateUserForm />
			</Main>
			<Footer />
		</Layout>
	);
}

export default CreateUser;

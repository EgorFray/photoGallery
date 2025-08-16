import Header from "../components/Header";
import Main from "../components/Main";
import NavBar from "../components/NavBar";
import CreatePostForm from "../components/CreatePostForm";
import Layout from "../components/Layout";
import Footer from "../components/Footer";

function CreatePost() {
	return (
		<Layout>
			<NavBar />
			<Header />
			<Main>
				<CreatePostForm />
			</Main>
			<Footer />
		</Layout>
	);
}

export default CreatePost;

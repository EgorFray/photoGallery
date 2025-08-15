import Header from "../components/Header";
import Main from "../components/Main";
import NavBar from "../components/NavBar";
import CreatePostForm from "../components/CreatePostForm";

function CreatePost() {
	return (
		<>
			<NavBar />
			<Header />
			<Main>
				<CreatePostForm />
			</Main>
		</>
	);
}

export default CreatePost;

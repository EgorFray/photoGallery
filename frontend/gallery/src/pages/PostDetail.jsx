import Footer from "../components/Footer";
import Main from "../components/Main";
import NavBar from "../components/NavBar";
import PostContent from "../components/PostContent";

function PostDetail() {
	return (
		<div>
			<NavBar />
			<Main>
				<PostContent />
			</Main>
			<Footer />
		</div>
	);
}

export default PostDetail;

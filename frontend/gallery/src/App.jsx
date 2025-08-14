import { PostsProvider } from "./context/PostsContext";
import PostDetail from "./pages/PostDetail";
import { Routes, Route } from "react-router-dom";
import MainPage from "./pages/MainPage";
import { PostDetailProvider } from "./context/PostDetailContext";
import CreatePost from "./pages/CreatePost";

function App() {
	return (
		<div>
			<PostsProvider>
				<PostDetailProvider>
					<Routes>
						<Route index element={<MainPage />} />
						<Route path="/create" element={<CreatePost />} />
						<Route path="posts/:id" element={<PostDetail />} />
					</Routes>
				</PostDetailProvider>
			</PostsProvider>
		</div>
	);
}

export default App;

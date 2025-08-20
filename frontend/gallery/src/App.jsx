import { PostsProvider } from "./context/PostsContext";
import PostDetail from "./pages/PostDetail";
import { Routes, Route } from "react-router-dom";
import MainPage from "./pages/MainPage";
import CreatePost from "./pages/CreatePost";
import PageNotFound from "./pages/PageNotFound";
import Homepage from "./pages/Homepage";
import Login from "./pages/Login";
import { AuthProvider } from "./context/FakeAuthContext";

function App() {
	return (
		<div>
			<AuthProvider>
				<PostsProvider>
					<Routes>
						<Route index element={<MainPage />} />
						<Route path="/homepage" element={<Homepage />} />
						<Route path="/login" element={<Login />} />
						<Route path="/create" element={<CreatePost />} />
						<Route path="posts/:id" element={<PostDetail />} />
						<Route path="*" element={<PageNotFound />} />
					</Routes>
				</PostsProvider>
			</AuthProvider>
		</div>
	);
}

export default App;

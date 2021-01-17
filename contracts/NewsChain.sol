pragma solidity ^0.8.0;

contract NewsChain {

    struct Writer {
        string name;
        string title;
        string email;
        int created_at;
    }

    struct Post {
        string title;
        string details;
        address writer;
        int created_at;
        int updated_at;
    }

    address owner;

    mapping(address => Writer) writers;
    mapping(string => Post) posts;
    mapping(address => bool) isWriterExists;
    mapping(string => bool) isPostIdExists;

    event PostCreated(string id, string title, int createdAt);
    event PostUpdated(string id, string title);

    function createAuthor(string memory name, string memory title, string memory email, int created_at) public payable {
        require(bytes(name).length > 0, "Name is required");
        require(bytes(title).length > 0, "Title is required");
        require(bytes(email).length > 0, "Email is required");
        require(created_at > 0, "CreatedAt is invalid");
        require(isWriterExists[msg.sender] == false, "User already registered");

        isWriterExists[msg.sender] = true;
        writers[msg.sender] = Writer(name, title, email, created_at);
    }

    function getAuthorProfile() public view returns (string memory name, string memory title, string memory email, int created_at){
        require(isWriterExists[msg.sender] == true, "Author is not registered");

        Writer memory author = writers[msg.sender];
        name = author.name;
        title = author.title;
        email = author.email;
        created_at = author.created_at;
    }

    function updateAuthor(string memory name, string memory title, string memory email) public payable {
        require(bytes(name).length > 0, "Name is required");
        require(bytes(title).length > 0, "Title is required");
        require(bytes(email).length > 0, "Email is required");
        require(isWriterExists[msg.sender] == true, "User isn't registered");

        writers[msg.sender].name = name;
        writers[msg.sender].title = title;
        writers[msg.sender].email = email;
    }

    function isAuthorRegistered() public view returns (bool){
        return isWriterExists[msg.sender];
    }

    function isOwner() public view returns (bool){
        return msg.sender == owner;
    }

    function createPost(string memory id, string memory title, string memory details, int created_at) public payable {
        require(isWriterExists[msg.sender] == true, "User not registered");
        require(isPostIdExists[id] == false, "PostId already exists");
        require(bytes(title).length > 0, "Title is required");
        require(bytes(details).length > 0, "Details is required");
        require(created_at > 0, "CreatedAt is invalid");

        isPostIdExists[id] = true;
        posts[id] = Post(title, details, msg.sender, created_at, 0);

        emit PostCreated(id, title, created_at);
    }

    function getPost(string memory id) public view returns (string memory title, address author, string memory details, int created_at, int updated_at) {
        require(isPostIdExists[id] == true, "PostId doesn't exists");

        Post memory post = posts[id];
        title = post.title;
        author = post.writer;
        details = post.details;
        created_at = post.created_at;
        updated_at = post.updated_at;
    }

    function updatePost(string memory id, string memory title, string memory details, int updated_at) public payable {
        require(isWriterExists[msg.sender] == true, "User not registered");
        require(isPostIdExists[id] == true, "PostId doesn't exists");
        require(bytes(title).length > 0, "Title is required");
        require(bytes(details).length > 0, "Details is required");
        require(updated_at > 0, "UpdatedAt is invalid");

        posts[id].title = title;
        posts[id].details = details;
        posts[id].updated_at = updated_at;

        emit PostUpdated(id, title);
    }
}

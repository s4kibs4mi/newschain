pragma solidity ^0.5.16;

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
        string tags;
        int created_at;
    }

    address owner;

    mapping(address => Writer) writers;
    mapping(string => Post) posts;
    mapping(address => bool) isWriterExists;
    mapping(string => bool) isPostIdExists;

    string portalName;

    function createAuthor(string memory name, string memory title, string memory email, int created_at) public payable {
        require(bytes(name).length > 0, "Name is required");
        require(bytes(title).length > 0, "Title is required");
        require(bytes(email).length > 0, "Email is required");
        require(created_at > 0, "CreatedAt is invalid");
        require(isWriterExists[msg.sender] == false, "User already registered");

        isWriterExists[msg.sender] = true;
        writers[msg.sender] = Writer(name, title, email, created_at);
    }

    function isAuthorRegistered() public returns (bool){
        return isWriterExists[msg.sender];
    }

    function isOwner() public returns (bool){
        return msg.sender == owner;
    }

    function createPost(string memory id, string memory title, string memory details, string memory tags, int created_at) public payable {
        require(isWriterExists[msg.sender] == true, "User not registered");
        require(isPostIdExists[id] == false, "PostId already exists");
        require(bytes(title).length > 0, "Title is required");
        require(bytes(details).length > 0, "Details is required");
        require(bytes(tags).length > 0, "Tags is required");
        require(created_at > 0, "CreatedAt is invalid");

        isPostIdExists[id] = true;
        posts[id] = Post(title, details, msg.sender, tags, created_at);
    }

    function getAuthorProfile() public returns (string memory name, string memory title, string memory email, int created_at){
        require(isWriterExists[msg.sender] == true, "Author is not registered");
        name = writers[msg.sender].name;
        title = writers[msg.sender].title;
        email = writers[msg.sender].email;
        created_at = writers[msg.sender].created_at;
    }
}

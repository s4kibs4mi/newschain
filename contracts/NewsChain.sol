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
        string[] tags;
        int created_at;
    }

    address owner;

    mapping(address => Writer) writers;
    mapping(address => Post) posts;
    mapping(address => bool) isWriterExists;

    string portalName;

    function createAuthor(string memory name, string memory title, string memory email, int created_at) public returns (bool){
        require(msg.sender == owner, "You don't have permission to create user");
        require(bytes(name).length > 0, "Name is required");
        require(bytes(title).length > 0, "Title is required");
        require(bytes(email).length > 0, "Email is required");
        require(created_at > 0, "CreatedAt is invalid");
        require(isWriterExists[msg.sender] == false, "User already registered");
        writers[msg.sender] = Writer(name, title, email, created_at);
        return true;
    }
}

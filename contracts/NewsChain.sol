pragma solidity ^0.5.16;

contract NewsChain {

    struct Writer {
        string name;
        string title;
        string password;
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
    mapping(string => bool) isWriterExists;

    string portalName;

    function createAuthor() public returns (bool){

        return false;
    }

    function createPost() public returns (bool) {

        return false;
    }
}

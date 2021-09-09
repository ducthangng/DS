#include <iostream>
#include <string>
#include <queue>
#include <unordered_map>
using namespace std;

#define EMPTY_STRING ""

struct Node
{
    char ch;
    int freq;
    Node *left, *right;
};

// Function to allocate a new tree node
Node *getNode(char ch, int freq, Node *left, Node *right)
{
    Node *node = new Node();
    node->ch = ch;
    node->freq = freq;
    node->left = left;
    node->right = right;
    return node;
}

//Comparison object to be used to order the heap
struct comp
{
    bool operator()(const Node *l, const Node *r) const
    {
        return l->freq > r->freq;
        // The highest priority item has the lowest frequency
    }
};

class HuffmanTree
{
private:
    unordered_map<char, int> freq;
    unordered_map<char, string> huffmanCode;
    priority_queue<Node *, vector<Node *>, comp> pq;

public:
    HuffmanTree(unordered_map<char, int> freq, unordered_map<char, string> huffmanCode)
    {
        this->freq = freq;
        this->huffmanCode = huffmanCode;
    }

    bool isLeaf(Node *root)
    {
        return root->left == nullptr && root->right == nullptr;
    }

    void encode(Node *root, string str)
    {
        if (root == nullptr)
        {
            return;
        }

        // Found a leaf node
        if (isLeaf(root))
        {
            huffmanCode[root->ch] = (str != EMPTY_STRING) ? str : "1";
        }

        encode(root->left, str + "0");
        encode(root->right, str + "1");
    }

    void decode(Node *root, int &index, string str)
    {
        if (root == nullptr)
        {
            return;
        }

        if (isLeaf(root))
        {
            cout << root->ch;
            return;
        }

        index++;

        if (str[index] == '0')
        {
            decode(root->left, index, str);
        }
        else
        {
            decode(root->right, index, str);
        }
    }

    void buildHuffmanTree(string text)
    {
        for (auto pair : freq)
        {
            pq.push(getNode(pair.first, pair.second, nullptr, nullptr));
        }

        // do till there is no more than one node in the queue
        while (pq.size() != 1)
        {
            // Remove two nodes of the highest priority
            // (the lowest frequency) from the queue
            Node *left = pq.top();
            pq.pop();
            Node *right = pq.top();
            pq.pop();

            int sum = left->freq + right->freq;
            pq.push(getNode('\0', sum, left, right));
        }

        // root stores pointer to the root of Huffman Tree
        Node *root = pq.top();

        // Traverse the Huffman Tree and store Huffman codes
        // in a map. Also, print them
        encode(root, EMPTY_STRING);

        cout << "Huffman Codes are: " << endl;
        for (auto pair : huffmanCode)
        {
            cout << pair.first << " " << pair.second << endl;
        }
    }

    void encodedstring(string text)
    {
        string str;
        for (char ch : text)
        {
            str += huffmanCode[ch];
        }
        cout << "The original text is: " << text << endl;
        cout << "The encoded text is: " << str << endl;
    }

    void decodedstring(string text)
    {
        Node *root = pq.top();
        cout << "The decoded text is: ";
        if (isLeaf(root))
        {
            while (root->freq--)
            {
                cout << root->ch;
            }
        }
        else
        {
            int index = -1;
            while (index < (int)text.size() - 1)
            {
                decode(root, index, text);
            }
        }
    }
};

int main()
{
    unordered_map<char, int> freq;
    unordered_map<char, string> hf;
    freq['a'] = 1;
    freq['b'] = 2;
    freq['c'] = 3;
    freq['d'] = 4;
    freq['e'] = 5;
    freq['f'] = 6;
    freq['g'] = 7;
    freq['h'] = 8;
    freq['i'] = 9;
    freq['j'] = 10;
    freq['k'] = 11;
    freq['l'] = 12;
    freq['m'] = 13;
    freq['n'] = 14;
    freq['o'] = 15;
    freq['p'] = 16;
    freq['q'] = 17;
    freq['r'] = 18;
    freq['s'] = 19;
    freq['t'] = 20;
    freq['u'] = 21;
    freq['v'] = 22;
    freq['w'] = 23;
    freq['x'] = 24;
    freq['y'] = 25;
    freq['z'] = 26;
    HuffmanTree huff(freq, hf);
    huff.buildHuffmanTree("nguyenducthang");
    huff.encodedstring("nguyenanhtai");
    huff.decodedstring("111011111100110110000100111101101111101110111111101011011111000101");
}

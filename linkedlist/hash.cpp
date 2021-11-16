#include <iostream> // not actually required for the hash
#include <string>

auto main() -> int
{
	const std::string input = "12";
	const std::hash<std::string> hasher;
	const auto hashResult = hasher(input);

	std::cout << "Input hash is: " << hashResult << std::endl;
}
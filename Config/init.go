package config

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
	Namespace  string
	fish       bool
	zsh        bool
	bash       bool
	Key        string
	SecretName string
	Kubeconfig string
	prefix     string
}

func NewConfig() *Config {
	var c Config
	var namespace, prefix string
	var Fish, Zsh, Bash bool
	flag.StringVar(&namespace, "n", "default", "Namespace du secret")
	flag.StringVar(&prefix, "p", "secret", "Programme nane fonction AutoCompletion")
	flag.BoolVar(&Fish, "fish", false, "Generate fish completion")
	flag.BoolVar(&Zsh, "zsh", false, "Generate zsh completion")
	flag.BoolVar(&Bash, "bash", false, "Generate bash completion")
	flag.Parse()

	c.Namespace = namespace
	c.fish = Fish
	c.zsh = Zsh
	c.bash = Bash
	c.prefix = prefix
	ret := c.AutoComplete()
	if ret {
		os.Exit(0)
	}

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: secret -n <namespace> secret-name key")
		os.Exit(1)
	}
	var key string
	if len(args) > 1 {
		key = args[1]
	}

	kubeconfig := os.Getenv("KUBECONFIG")
	if kubeconfig == "" {
		fmt.Println("KUBECONFIG not found")
		os.Exit(1)
	}

	return &Config{
		Namespace:  namespace,
		fish:       Fish,
		zsh:        Zsh,
		bash:       Bash,
		Key:        key,
		SecretName: args[0],
		Kubeconfig: kubeconfig,
	}
}

func (c *Config) AutoComplete() bool {
	if c.fish {
		fish(c.prefix)
		return true
	}
	if c.zsh {
		fmt.Println("Zsh completion enabled")
		// TODO generate zsh completion
		return true
	}
	if c.bash {
		fmt.Println("Bash completion enabled")
		// TODO generate bash completion
		return true
	}
	return false
}

func fish(prefix string) {
	fishCompletion := fmt.Sprintf(`
function __fish_secret_using_n
    set -l cmd (commandline -opc)
    for i in (seq (count $cmd))
        if test $cmd[$i] = "-n"; or test $cmd[$i] = "--namespace"
            if test (count $cmd) -eq $i
                return 0
            end
        end
    end
    return 1
end

function __fish_secret_namespaces
    kubectl get namespaces -o name | string replace "namespace/" "" 2>/dev/null
end


function __fish_secret_current_namespace
    set -l cmd (commandline -opc)
    for i in (seq (count $cmd))
        if test $cmd[$i] = "-n"; or test $cmd[$i] = "--namespace"
            if test (count $cmd) -gt $i
                echo $cmd[(math $i + 1)]
                return
            end
        end
    end
    # Fallback: default namespace
    echo "default"
end

function __fish_secret_list_names
    set -l ns (__fish_secret_current_namespace)
    kubectl get secrets -n $ns -o name | string replace "secret/" "" 2>/dev/null
end

function __fish_secret_current_namespace
    set -l cmd (commandline -opc)
    for i in (seq (count $cmd))
        if test $cmd[$i] = "-n"; or test $cmd[$i] = "--namespace"
            if test (count $cmd) -gt $i
                echo $cmd[(math $i + 1)]
                return
            end
        end
    end
    # Fallback: default namespace
    echo "default"
end


function __fish_secret_current_secret
    set -l cmd (commandline -opc)
    set -l found_non_option 0
    
    for i in (seq (count $cmd))
        if not string match -qr '^-' -- $cmd[$i]
            set found_non_option (math $found_non_option + 1)
            
            if test $found_non_option -eq 3
                echo $cmd[ $i ]
                return
            end
        end
    end
end

function __fish_secret_keys
    set -l ns (__fish_secret_current_namespace)
    set -l secret (__fish_secret_current_secret)
    if test -n "$ns" -a -n "$secret"
        echo "$secret $ns" > debug.log
        kubectl get secret $secret -n $ns -o json 2>/dev/null | jq -r '.data | keys[]'
    end
end

complete -c %s -e
complete -c %s -f -n "__fish_secret_using_n" -a "(__fish_secret_namespaces)" -d "Namespace"
complete -c %s -f -s n -l namespace -d "Namespace Kubernetes"
complete -c %s -f -n '__fish_seen_subcommand_from -n --namespace' -a "(__fish_secret_list_names)" -d "Secret name"
complete -c %s -f -n '__fish_secret_keys' -a "(__fish_secret_keys)" -d "Secret key"
		`, prefix, prefix, prefix, prefix, prefix)
	fmt.Println(fishCompletion)
}
